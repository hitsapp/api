import type { NextPage } from "next";
import { Leaderboard, Nav } from "../components";

const Home: NextPage = () => {
  return (
    <>
      <Nav />
      <Leaderboard />
    </>
  );
};

export default Home;
