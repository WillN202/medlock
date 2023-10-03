import 'tailwindcss/tailwind.css'
import {parseCookies} from 'nookies'
import auth from "../app/user/auth";

const {AuthContext} = auth;

function SafeHydrate({children}) {
  return (
    <div suppressHydrationWarning>
      {typeof window === 'undefined' ? null : children}
    </div>
  )
}

const CookieAuthContext = ({children}) => {
  const cookies = parseCookies();
  const account = cookies.account;
  const user = auth.toUser(account);
  return <AuthContext initialUser={user}>{children}</AuthContext>
}

function MyApp({Component, pageProps}) {
  return (
    <SafeHydrate>
      <CookieAuthContext>
        <Component {...pageProps} />
      </CookieAuthContext>
    </SafeHydrate>
  )
}

export default MyApp
