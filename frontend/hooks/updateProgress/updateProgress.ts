import axios from "axios";
import {Progress} from "../../app/achievement/achievement";

const updateProgress = async (p: Progress, id: String, achievement: String): Promise<Boolean> => {
  try {
    await axios.put(`http://localhost:4000/students/${id}/achievements/${achievement}/progress`,
      {
        "progress": p
      }
    )
  } catch (e) {
    console.log(e)
    return false
  }
  return true
}

export default {updateProgress}