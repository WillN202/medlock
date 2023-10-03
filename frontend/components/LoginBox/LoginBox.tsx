import {useState} from "react";
import hooks from "../../hooks/useLogin/useLogin";
import {useRouter} from "next/router";
import cookieCutter from "cookie-cutter";
import auth from "../../app/user/auth";

const LoginBox = () => {

  const [code, setCode] = useState("")
  const [showError, setShowError] = useState(false)
  const router = useRouter()

  const updateCode = (e) => {
    let c = e.currentTarget.value.toUpperCase()
    setCode(c.substring(0, 4))
  }

  const login = () => {
    const a = hooks.useLogin(code)
    a.then(account => {
      if (!account) {
        setShowError(true)
        return
      }
      cookieCutter.set("account", auth.toCookie(account))
      router.push("/dashboard")
    })
  }

  const disabled = code.length !== 4;
  return (
    <div className={"flex flex-col justify-center border text-center p-10 rounded bg-gray-200 text-3xl"}>
      <h1 className={"mt-5 mb-5"}>Login with your code</h1>
      {showError &&
        <p className={"text-lg bg-red-200 mb-3 p-2"}>Could not log in. Have you entered the correct code?</p>}
      <div>
        <input className={"text-center p-2 rounded-lg w-32"} inputMode={"text"} size={4} value={code} onChange={updateCode}
               type="text"
               placeholder="Code..."/>
      </div>
      <div>
        <button onClick={() => login()} className={`${disabled ? "bg-gray-400 text-gray-500" : "bg-blue-400 hover:bg-blue-500"} rounded-lg my-5 p-2 px-5`}
                disabled={disabled}>Login
        </button>
      </div>
    </div>
  )
};

export default LoginBox;