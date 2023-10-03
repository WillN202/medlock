import api from "../../hooks/useAchievements/useAchievements";
import {useEffect, useState} from "react";
import {Achievement} from "../../app/achievement/achievement";
import auth from "../../app/user/auth";
import AchievementBlock from "../AchievementBlock/AchievementBlock";

const {useRequireAuth, OnlyStudent} = auth;

const render = () => {
  const user = useRequireAuth();
  const [achievements, setAchievements] = useState<Achievement[]>([])
  useEffect(() => {
    if (user.isTeacher()) {
      return
    }
    const resp = api.useAchievements(user.id)
    resp.then(achvs => {
      setAchievements(achvs)
    })
  }, [user.id])

  return (
    <>
      <OnlyStudent>
        <div className={"grid grid-cols-4 md:grid-cols-3 sm:grid-cols-1 gap-4"}>
          {achievements && achievements.map(a => <div key={a.name} data-testid={`${a.name}_block`}>
              <AchievementBlock achievement={a}/>
            </div>
          )}
        </div>
      </OnlyStudent>
    </>
  );
}

export default render