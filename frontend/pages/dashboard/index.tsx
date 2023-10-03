import Head from "next/head"
import auth from "../../app/user/auth"
import AchievementList from "../../components/AchievementList/AchievementList";
import Header from "../../components/Header/Header";

const {OnlyAuth} = auth

export default function Home() {
  return (
    <>
      <Header/>
      <OnlyAuth redirectTo={"/"}>
        <div className="flex flex-col items-center justify-center min-h-screen py-2">
          <Head>
            <title>Academic Achievements - Medlock Primary School</title>
            <link rel="icon" href="/favicon.ico"/>
          </Head>
          <div className="max-w-5xl text-center">
            <h1 className="text-3xl mb-10">My Achievement Progress</h1>
            <AchievementList/>
          </div>
        </div>
      </OnlyAuth>
    </>
  )
}
