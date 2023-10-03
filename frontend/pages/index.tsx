import Head from "next/head"
import LoginBox from "../components/LoginBox/LoginBox"
import auth from "../app/user/auth"

const {OnlyUnauth} = auth

export default function Home() {
  return (
    <OnlyUnauth redirectTo={"/dashboard"}>
      <div className="flex flex-col items-center justify-center min-h-screen py-2">
        <Head>
          <title>Academic Achievements - Medlock Primary School</title>
          <link rel="icon" href="/favicon.ico"/>
        </Head>

        <LoginBox/>
      </div>
    </OnlyUnauth>
  )
}
