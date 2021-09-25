import type { NextPage } from "next";
import { Leaderboard, Nav, Subtitle, Title } from "../components";
import styled from "styled-components";

const Container = styled.div`
  width: 80%;
  height: 80vh;
  max-width: 1200px;
  margin: 0 auto;
`;

const MainContent = styled.div`
  display: flex;
  height: 100%;
  flex-direction: row;
`;

const RightContainer = styled.div`
  display: flex;
  flex-direction: column;
  margin-left: 20px;
  width: 100%;
`;

const RightFlexContainers = styled.div`
  margin: 0;
  padding: 24px 24px;
  flex: 1 1;
  border-radius: 8px;
  background: ${({ theme }) => theme.layoutDark};
  box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.25);

  &:last-of-type {
    margin-top: 20px;
  }
`;

const Home: NextPage = () => {
  return (
    <Container>
      <Nav />
      <MainContent>
        <Leaderboard />
        <RightContainer>
          <RightFlexContainers>
            <Title>Create a Hit</Title>
            <Subtitle>Generate a SVG Image for your site</Subtitle>
          </RightFlexContainers>
          <RightFlexContainers>
            <Title>Copy URL</Title>
            <Subtitle>Choose from the following:</Subtitle>
          </RightFlexContainers>
        </RightContainer>
      </MainContent>
    </Container>
  );
};

export default Home;
