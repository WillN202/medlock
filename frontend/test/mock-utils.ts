import {User, UserType} from "../app/user/auth";

type MockUserOpts = {
  id?: string
  type?: UserType
  name?: string
}

export const mockUser = (opts?: MockUserOpts): User => ({
  id: opts?.id || "p1000",
  type: opts?.type || "Student",
  name: opts?.name || "mock user",
  isTeacher(): boolean {
    return false
  },
  isStudent(): boolean {
    return true
  }
});
