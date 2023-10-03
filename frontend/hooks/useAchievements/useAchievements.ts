import {Achievement} from "../../app/achievement/achievement";
import axios from "axios";

const useAchievements = async (studentId: string): Promise<Achievement[] | undefined> => {
  let resp: any
  try {
    resp = await axios.get(`http://localhost:4000/students/${studentId}/achievements`)
  } catch (e) {
    console.log(e)
    return undefined
  }
  return resp.data.achievements.map(a => ({
    id: a.achievement.id,
    name: a.achievement.name,
    progress: a.progress,
  }))
}

export default {useAchievements}