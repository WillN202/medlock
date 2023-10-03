import * as React from "react";
import {fireEvent, render} from "@testing-library/react";
import {MemoryRouter} from "next-router-mock";
import {RouterContext} from "next/dist/shared/lib/router-context";
import auth from "./auth";
import {mockUser} from "../../test/mock-utils";

const {AuthContext, OnlyAuth, OnlyUnauth, useAuthContext, useRequireAuth} = auth;

describe("AuthContext", () => {

  const TestAuthConsumer = () => {
    const ctx = useAuthContext();
    if (!ctx.isAuth()) {
      return <>
        <p>user not set</p>
      </>;
    }
    return (<>
      <p>{ctx.user?.id}</p>
      <button type="button" onClick={ctx.signOut}>sign out</button>
    </>);
  };

  it("should return null if used without context", async () => {
    const screen = render(<TestAuthConsumer/>);
    expect(await screen.findByText("user not set")).toBeInTheDocument();
  });

  it('should indicate where user is not set', async () => {
    const user = undefined;
    const screen = render(
        <AuthContext initialUser={user}>
          <TestAuthConsumer/>
        </AuthContext>);
    expect(await screen.findByText("user not set")).toBeInTheDocument();
  });

  it("should return user if set in context", async () => {
    const user = mockUser({id: "p001"});
    const screen = render(
        <AuthContext initialUser={user}>
          <TestAuthConsumer/>
        </AuthContext>);
    expect(await screen.queryByText("p001")).toBeInTheDocument();
    expect(await screen.queryByText("user not set")).not.toBeInTheDocument();
  });

  it("should sign out", async () => {
    const user = mockUser({id: "p001"});
    const screen = render(
        <AuthContext initialUser={user}>
          <TestAuthConsumer/>
        </AuthContext>);
    const button = screen.getByText(/sign out/);
    fireEvent.click(button);
    expect(await screen.findByText("user not set")).toBeInTheDocument();
  });

});

describe("useRequireAuthContext", () => {

  const ComponentRequiringUser = () => {
    const user = useRequireAuth();
    return <p>{user.id}</p>;
  };

  afterEach(jest.restoreAllMocks);

  it("should return user", () => {
    const user = mockUser({id: "p00"});
    const screen = render(<AuthContext initialUser={user}><ComponentRequiringUser/></AuthContext>);
    expect(screen.baseElement.textContent).toEqual("p00");
  });

  it('should throw error if there is no user', () => {
    jest.spyOn(console, 'error').mockImplementation(() => {});
    expect(() => render(<AuthContext initialUser={undefined}>
        <ComponentRequiringUser/>
    </AuthContext>))
    .toThrow("user missing from AuthContext");
  });

});

describe("OnlyAuth", () => {

  it("should not render children if unauth", async () => {
    const router = new MemoryRouter("/dashboard");
    const screen = render(
        <RouterContext.Provider value={router}>
          <AuthContext initialUser={undefined}>
            <OnlyAuth><p>allowed in</p></OnlyAuth>
          </AuthContext>
        </RouterContext.Provider>
    );
    expect(await screen.queryByText("allowed in")).not.toBeInTheDocument();
    expect(router.pathname).toEqual("/dashboard");
  });

  it("should redirect if unauthorized", async () => {
    const router = new MemoryRouter("/dashboard");
    const screen = render(
        <RouterContext.Provider value={router}>
          <AuthContext initialUser={undefined}>
            <OnlyAuth redirectTo="/unauth"><p>allowed in</p></OnlyAuth>
          </AuthContext>
        </RouterContext.Provider>
    );
    expect(await screen.queryByText("allowed in")).not.toBeInTheDocument();
    expect(router.pathname).toEqual("/unauth");
  });

  it("should render children if auth", async () => {
    const router = new MemoryRouter("/dashboard");
    const user = mockUser({id: "p001"});
    const screen = render(
        <RouterContext.Provider value={router}>
          <AuthContext initialUser={user}>
            <OnlyAuth redirectTo="/unauth"><p>allowed in</p></OnlyAuth>
          </AuthContext>
        </RouterContext.Provider>
    );
    expect(await screen.queryByText("allowed in")).toBeInTheDocument();
    expect(router.pathname).toEqual("/dashboard");
  });

});

describe("OnlyUnauth", () => {

  const user = mockUser({id: "p000001"});

  it("should not render children if auth", async () => {
    const router = new MemoryRouter("/unauth");
    const screen = render(
        <RouterContext.Provider value={router}>
          <AuthContext initialUser={user}>
            <OnlyUnauth><p>children</p></OnlyUnauth>
          </AuthContext>
        </RouterContext.Provider>
    );
    expect(await screen.queryByText("children")).not.toBeInTheDocument();
    expect(router.pathname).toEqual("/unauth");
  });

  it("should redirect if auth", async () => {
    const router = new MemoryRouter("/unauth");
    const screen = render(
        <RouterContext.Provider value={router}>
          <AuthContext initialUser={user}>
            <OnlyUnauth redirectTo="/dashboard"><p>unauth page</p></OnlyUnauth>
          </AuthContext>
        </RouterContext.Provider>
    );
    expect(await screen.queryByText("unauth page")).not.toBeInTheDocument();
    expect(router.pathname).toEqual("/dashboard");
  });

  it("should render if unauth", async () => {
    const router = new MemoryRouter("/unauth");
    const screen = render(
        <RouterContext.Provider value={router}>
          <AuthContext initialUser={undefined}>
            <OnlyUnauth redirectTo="/dashboard"><p>unauth page</p></OnlyUnauth>
          </AuthContext>
        </RouterContext.Provider>
    );
    expect(await screen.queryByText("unauth page")).toBeInTheDocument();
    expect(router.pathname).toEqual("/unauth");
  });

});
