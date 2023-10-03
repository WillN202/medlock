import {useRouter} from "next/router";
import React, {createContext, useContext, useEffect, useState} from "react";

export type UserType = "Teacher" | "Student"

export const IsTeacher = (userType) => userType === "Teacher"
export const IsStudent = (userType) => userType === "Student"

export type User = {
  id: string
  name: string
  type: UserType
  isTeacher(): boolean
  isStudent(): boolean
}

export type UserContext = {
  user: User | undefined
  signOut: () => void
  isAuth: () => boolean
  isTeacher(): boolean
  isStudent(): boolean
}

const defaultUserContext = {
  user: undefined,
  isAuth: () => false,
  isTeacher: () => false,
  isStudent: () => false,
  signOut: () => {
  },
};

const AuthHolder = createContext<UserContext>(defaultUserContext);

export interface AuthContextProps {
  children: React.ReactNode
  initialUser: User | undefined
}

function clearCookie(name: string) {
  document.cookie = `${name}=;Path=/;Expires=Thu, 01 Jan 1970 00:00:01 GMT;`;
}

const AuthContext = ({children, initialUser}: AuthContextProps) => {
  const [user, setUser] = useState<User | undefined>(initialUser);

  useEffect(() => {
    setUser(initialUser)
  }, [initialUser]);

  const signOut = () => {
    setUser(undefined);
    clearCookie("account");
  };
  const isAuth = (): boolean => !!user;
  const isTeacher = (): boolean => !!user && user.isTeacher();
  const isStudent = (): boolean => !!user && user.isStudent();
  return <AuthHolder.Provider value={{user, isAuth, signOut, isTeacher, isStudent}}>
    {children}
  </AuthHolder.Provider>;
};

interface MustAuthProps {
  children: React.ReactNode
  redirectTo?: string
}

const useAuthContext = (): UserContext => {
  const userContext = useContext(AuthHolder);
  if (!userContext) {
    console.warn("user context is undefined, only use within AuthContext");
    return defaultUserContext;
  }
  return userContext;
};

const useRequireAuth = (): User => {
  const userContext = useContext(AuthHolder);
  if (!userContext || !userContext.isAuth()) {
    throw new Error("user missing from AuthContext");
  }
  return userContext.user as User;
};

const OnlyAuth = ({children, redirectTo}: MustAuthProps) => {
  const userContext = useAuthContext();
  const router = useRouter();

  const userAuthed = userContext && userContext.isAuth();

  useEffect(() => {
    if (!userAuthed && redirectTo)
      void router.push(redirectTo);
  });

  return userAuthed ? <>{children}</> : null;
};

const OnlyUnauth = ({children, redirectTo}: MustAuthProps) => {
  const userContext = useAuthContext();
  const router = useRouter();

  const userAuthed = userContext && userContext.isAuth();

  useEffect(() => {
    if (userAuthed && redirectTo)
      void router.push(redirectTo);
  });

  return !userAuthed ? <>{children}</> : null;
};

const OnlyStudent = ({children, redirectTo}: MustAuthProps) => {
  const userContext = useAuthContext();
  const router = useRouter();

  const userAuthed = userContext && userContext.isAuth();

  useEffect(() => {
    if ((!userAuthed || userContext.isTeacher()) && redirectTo)
      void router.push(redirectTo);
  });

  return userContext.isStudent() ? <>{children}</> : null;
};

const OnlyTeacher = ({children, redirectTo}: MustAuthProps) => {
  const userContext = useAuthContext();
  const router = useRouter();

  const userAuthed = userContext && userContext.isAuth();

  useEffect(() => {
    if ((!userAuthed || userContext.isStudent()) && redirectTo)
      void router.push(redirectTo);
  });

  return userContext.isTeacher() ? <>{children}</> : null;
};

const toUser = (s: string): User | undefined => {
  if (!s) {
    return undefined
  }
  const split = s.split("#")
  if (split.length !== 3) {
    throw new Error("invalid user cookie")
  }
  let userType = split[2];
  if (userType !== "Student" && userType !== "Teacher") {
    throw new Error(`invalid user type ${userType}`)
  }
  return {
    isStudent(): boolean {
      return userType === "Student";
    }, isTeacher(): boolean {
      return userType === "Teacher";
    },
    id: split[0],
    name: split[1],
    type: userType
  }
}

const toCookie = (user: User): string => {
  return `${user.id}#${user.name}#${user.type}`
}

export default {
  AuthContext,
  useAuthContext,
  useRequireAuth,
  OnlyAuth,
  OnlyUnauth,
  OnlyTeacher,
  OnlyStudent,
  toCookie,
  toUser,
};