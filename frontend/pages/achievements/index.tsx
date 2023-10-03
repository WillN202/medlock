import Head from "next/head"
import auth from "../../app/user/auth"
import {useState} from "react";

const {OnlyAuth} = auth

const defaultAchvs = [
  "Set Up Recycling Boxes",
  "Bring in Eco Friendly Water Bottle",
  "Fix a Broken Toy",
  "Stitch up a Hole in some Clothing",
  "Start a Compost Heap",
  "Plant some Seeds",
  "Recycle 10 Batteries",
  "Make a Sculpture out of Bottle Caps",
  "Donate Old Clothing",
];

export default function Home() {


  return (
    <OnlyAuth redirectTo={"/"}>
      <Head>
        <title>Academic Achievements - Medlock Primary School</title>
        <link rel="icon" href="/favicon.ico"/>
      </Head>
      <div className="flex flex-col items-center justify-center min-h-screen py-2">
        <div className="flex flex-col min-w-[50%] text-center">
          <h1 className="text-3xl mb-10">All Achievements</h1>
          <AllAchievementsList/>
        </div>
      </div>
    </OnlyAuth>
  )
}

const generateNumbers = () => {
  const completed = Math.floor(Math.random() * 25);
  const inprogress = Math.floor(Math.random() * (30 - completed));
  const notstarted = 30 - completed - inprogress;
  return {completed, inprogress, notstarted};
}

const initialNumbers = () => {
  let numbers = [];
  for (let i = 0; i < defaultAchvs.length; i++) {
    numbers.push(generateNumbers());
  }
  return numbers;
}

const AllAchievementsList = () => {
  const [numbers, setNumbers] = useState(initialNumbers)
  const [achvs, setAchvs] = useState<string[]>(defaultAchvs);
  const [value, setValue] = useState<string>("");

  const handleAddNewAchievement = () => {
    setAchvs([...achvs, value]);
  }

  return <div>
    {achvs.map((a, i) => {
      const {completed, inprogress, notstarted} = numbers[i % numbers.length];
      return <div
        className="p-3 flex text-left text-xl border my-1 w-full bg-gray-100 justify-between hover:bg-gray-200">
        <a href={`/achievements/view/${a.replaceAll(" ", "-")}`}>{a}</a>
        <div className="flex gap-2 w-44 justify-evenly items-center align-middle">
          <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" viewBox="0 0 20 20" fill="green">
            <path
              d="M13 6a3 3 0 11-6 0 3 3 0 016 0zM18 8a2 2 0 11-4 0 2 2 0 014 0zM14 15a4 4 0 00-8 0v3h8v-3zM6 8a2 2 0 11-4 0 2 2 0 014 0zM16 18v-3a5.972 5.972 0 00-.75-2.906A3.005 3.005 0 0119 15v3h-3zM4.75 12.094A5.973 5.973 0 004 15v3H1v-3a3 3 0 013.75-2.906z"/>
          </svg>
          {completed}
          <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" viewBox="0 0 20 20" fill="orange">
            <path
              d="M13 6a3 3 0 11-6 0 3 3 0 016 0zM18 8a2 2 0 11-4 0 2 2 0 014 0zM14 15a4 4 0 00-8 0v3h8v-3zM6 8a2 2 0 11-4 0 2 2 0 014 0zM16 18v-3a5.972 5.972 0 00-.75-2.906A3.005 3.005 0 0119 15v3h-3zM4.75 12.094A5.973 5.973 0 004 15v3H1v-3a3 3 0 013.75-2.906z"/>
          </svg>
          {inprogress}
          <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" viewBox="0 0 20 20" fill="red">
            <path
              d="M13 6a3 3 0 11-6 0 3 3 0 016 0zM18 8a2 2 0 11-4 0 2 2 0 014 0zM14 15a4 4 0 00-8 0v3h8v-3zM6 8a2 2 0 11-4 0 2 2 0 014 0zM16 18v-3a5.972 5.972 0 00-.75-2.906A3.005 3.005 0 0119 15v3h-3zM4.75 12.094A5.973 5.973 0 004 15v3H1v-3a3 3 0 013.75-2.906z"/>
          </svg>
          {notstarted}
        </div>
      </div>
    })}
    <div className="p-3 flex text-left text-xl border my-1 w-full bg-gray-100 justify-between hover:bg-gray-200">
      <input className="bg-gray-100" type="text" placeholder="New Achievement..." size={80} onChange={(e) => setValue(e.target.value)} value={value}/>
      <div className="flex items-center w-5 h-5 bg-gray-200 hover:bg-gray-300 rounded-full">
        <button onClick={handleAddNewAchievement} disabled={!value.length}>
          <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" viewBox="0 0 20 20" fill="green">
            <path fillRule="evenodd"
                  d="M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z"
                  clipRule="evenodd"/>
          </svg>
        </button>
      </div>
    </div>
  </div>;
}
