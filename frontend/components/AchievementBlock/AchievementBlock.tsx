import {Achievement, Progress} from "../../app/achievement/achievement";
import {useState} from "react";
import api from "../../hooks/updateProgress/updateProgress";
import auth from "../../app/user/auth";

interface AchievementBlockParams {
  achievement: Achievement
}

interface ProgressParams {
  progress: Progress
  id: String
  achievement: String
}

const ProgressCmp = ({progress, id, achievement}: ProgressParams) => {
  const [p, setP] = useState<Progress>(progress)

  const handleUpdate = async (p: Progress) => {
    const callSucceed = await api.updateProgress(p, id, achievement)
    if (!callSucceed) {
      return
    }
    setP(p)
  }

  return <div className={"flex justify-between px-2 py-2 m-2 bg-gray-500 rounded-lg max-w-min"}>
    <div className={`w-5 m-1 p-5 rounded-full ${p === "" ? "bg-red-500" : "bg-red-400/50"} hover:bg-red-500`} onClick={() => handleUpdate("")}/>
    <div className={`w-5 m-1 p-5 rounded-full ${p === "STARTED" ? "bg-yellow-400" : "bg-yellow-400/50"} hover:bg-yellow-400`} onClick={() => handleUpdate("STARTED")}/>
    <div className={`w-5 m-1 p-5 rounded-full ${p === "FINISHED" ? "bg-green-400" : "bg-green-600/50"} hover:bg-green-400`} onClick={() => handleUpdate("FINISHED")}/>
  </div>
}

function getRandomInt(max) {
  return Math.floor(Math.random() * max);
}

const colours = [
  "bg-blue-200",
  "bg-pink-200",
  "bg-purple-200",
]

const render = ({achievement}: AchievementBlockParams) => {
  
  const {id} = auth.useRequireAuth()

  return <div className={`grid grid-cols-1 content-between min-h-full justify-center content-between p-5 rounded-lg text-center h-60 ${colours[getRandomInt(3)]}`}>
    <h2 className={"text-2xl"}>{achievement.name}</h2>
    <div className={"flex min-w-full justify-center"}><ProgressCmp progress={achievement.progress} achievement={achievement.id} id={id}/></div>
  </div>;
}

export default render