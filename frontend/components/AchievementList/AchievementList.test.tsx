import {render} from "@testing-library/react";
import AchievementList from "./AchievementList";
import {RouterContext} from "next/dist/shared/lib/router-context";
import * as React from "react";
import {MemoryRouter} from "next-router-mock";
import {mockUser} from "../../test/mock-utils";
import auth from "../../app/user/auth";
import hooks from "../../hooks/useAchievements/useAchievements";
import {Achievement} from "../../app/achievement/achievement";

const {AuthContext} = auth;

describe("Achievement List", () => {
  it("should fetch achievements for id in auth context", async() => {
    jest.spyOn(hooks, "useAchievements").mockImplementation((): Promise<Achievement[] | undefined> =>
      Promise.resolve([{
        name: "ach001",
        progress: "STARTED",
        studentId: "stu001"
      }])
    );
    const router = new MemoryRouter("/dashboard");
    const user = mockUser();
    const screen = render(
      <RouterContext.Provider value={router}>
        <AuthContext initialUser={user}>
          <AchievementList/>
        </AuthContext>
      </RouterContext.Provider>
    )
    expect(await screen.findByTestId("ach001_block")).toBeInTheDocument()
    expect(await screen.findByText("ach001")).toBeInTheDocument()
  })
})