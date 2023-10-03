import {fireEvent, render, waitFor, within} from "@testing-library/react";
import LoginBox from "./LoginBox";
import {User} from "../../app/user/auth";
import hooks from "../../hooks/useLogin/useLogin";
import {MemoryRouter} from "next-router-mock";
import {RouterContext} from "next/dist/shared/lib/router-context";
import cookieCutter from "cookie-cutter";

describe("Login Box", () => {

  it("should have an input box with placeholder text for the code", () => {
    const screen = render(<LoginBox/>);

    const input = screen.getByRole("textbox") as HTMLInputElement;
    expect(input).toBeInTheDocument();

    expect(input.getAttribute("placeholder")).toEqual("Code...");
    expect(input.value).toEqual("");
  });

  it("should have a login button to submit a code", () => {
    const screen = render(<LoginBox/>);

    const input = screen.getByRole("button") as HTMLButtonElement;
    expect(input).toBeInTheDocument();

    const {getByText} = within(input);
    expect(getByText("Login")).toBeInTheDocument();
  });

  it("should make all letters uppercase", () => {
    const screen = render(<LoginBox/>);

    const input = screen.getByRole("textbox") as HTMLInputElement;
    expect(input).toBeInTheDocument();

    fireEvent.change(input, {target: {value: 'abcd'}})
    expect(input.value).toBe("ABCD");
  });

  it("should have a limit of 4 characters", () => {
    const screen = render(<LoginBox/>);

    const input = screen.getByRole("textbox") as HTMLInputElement;
    expect(input).toBeInTheDocument();

    fireEvent.change(input, {target: {value: 'abcdef'}})
    expect(input.value).toBe("ABCD");
  });

  it("should disable the button when code is less than 4 characters", () => {
    const screen = render(<LoginBox/>);

    const input = screen.getByRole("textbox") as HTMLInputElement;
    expect(input).toBeInTheDocument();
    const button = screen.getByRole("button") as HTMLButtonElement;
    expect(button).toBeInTheDocument();

    fireEvent.change(input, {target: {value: 'A'}})
    expect(button.disabled).toEqual(true);
    fireEvent.change(input, {target: {value: 'AB'}})
    expect(button.disabled).toEqual(true);
    fireEvent.change(input, {target: {value: 'ABC'}})
    expect(button.disabled).toEqual(true);
    fireEvent.change(input, {target: {value: 'ABCD'}})
    expect(button.disabled).toEqual(false);
  });

  it("should show an error message when login fails", async () => {
    jest.spyOn(hooks, "useLogin").mockImplementation((): Promise<User | undefined> =>
      Promise.resolve(undefined))
    const router = new MemoryRouter("/");
    const screen = render(
      <RouterContext.Provider value={router}>
        <LoginBox/>
      </RouterContext.Provider>
    );

    const input = screen.getByRole("textbox") as HTMLInputElement;
    expect(input).toBeInTheDocument();
    const button = screen.getByRole("button") as HTMLButtonElement;
    expect(button).toBeInTheDocument();

    expect(await screen.queryByText("Could not log in. Have you entered the correct code?")).not.toBeInTheDocument()
    fireEvent.change(input, {target: {value: 'ABCD'}});
    fireEvent.click(button);
    await waitFor(() => expect(hooks.useLogin).toHaveBeenCalled())
    expect(await screen.getByText("Could not log in. Have you entered the correct code?")).toBeInTheDocument()
    expect(router.pathname).toEqual("/");
  });


  it("should set account cookie and redirect to dashboard when login is successful", async () => {
    jest.spyOn(hooks, "useLogin").mockImplementation((): Promise<User | undefined> =>
      Promise.resolve({
        type: "Student",
        id: "123",
        name: "Mock Student"
      })
    );
    const setCookies = jest.spyOn(cookieCutter, "set")
    const router = new MemoryRouter("/");
    const screen = render(
      <RouterContext.Provider value={router}>
        <LoginBox/>
      </RouterContext.Provider>
    );

    const input = screen.getByRole("textbox") as HTMLInputElement;
    expect(input).toBeInTheDocument();
    const button = screen.getByRole("button") as HTMLButtonElement;
    expect(button).toBeInTheDocument();

    fireEvent.change(input, {target: {value: 'ABCD'}});
    fireEvent.click(button);
    await waitFor(() => expect(hooks.useLogin).toHaveBeenCalled());
    await waitFor(() => expect(setCookies).toHaveBeenCalled());
    expect(router.pathname).toEqual("/dashboard");
  });

});