import type { NextPage } from "next";
import { Leaderboard, Nav } from "../components";
import styled from "styled-components";

const Container = styled.div`
  width: 80%;
  max-width: 1200px;
  margin: 0 auto;
`;

const Home: NextPage = () => {
  return (
    <Container>
      <Nav />
      <Leaderboard />
    </Container>
  );
};

export default Home;
