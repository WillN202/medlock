import auth from "../../app/user/auth"

const {useAuthContext, OnlyAuth} = auth

const Header = () => {

  const ctx = useAuthContext()

  return (
    <div className="w-screen bg-gray-100 h-14 flex justify-end py-2 px-20">
      <OnlyAuth>
        <button className="bg-blue-300 rounded-full py-2 px-5" onClick={ctx.signOut}>
          Sign Out
        </button>
      </OnlyAuth>
    </div>
  );
}

export default Header