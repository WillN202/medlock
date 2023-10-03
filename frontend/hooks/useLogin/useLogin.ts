import axios from "axios";
import {User, IsStudent, IsTeacher} from "../../app/user/auth";

const useLogin = async (code: string): Promise<User | undefined> => {
  let resp: any
  try {
    resp = await axios.post("http://localhost:4000/login", {code})
  } catch (e) {
    console.log(e)
    return undefined
  }
  return {
    isStudent(): boolean {
      return IsStudent(resp.data.type);
    }, isTeacher(): boolean {
      return IsTeacher(resp.data.type);
    },
    id: resp.data.id,
    name: resp.data.name,
    type: resp.data.type,
  }
}

export default {useLogin}