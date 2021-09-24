import { useLeaderboard } from "../hooks/useLeaderboard";
import styled from "styled-components";
import { Subtitle, Title } from ".";

const Container = styled.div`
  padding: 24px 24px;
  max-width: 410px;
  background: ${({ theme }) => theme.layoutDark};
  border-radius: 8px;
  overflow-y: auto;
  overflow-x: hidden;
  box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.25);
`;

const List = styled.ul`
  list-style: none;
  margin: 0;
  padding: 0;
`;

const ListItem = styled.li`
  display: flex;
  flex-direction: row;
  align-items: center;
  overflow: hidden;
  white-space: nowrap;
`;

const Place = styled.span`
  min-width: 22px;
  font-size: 1.2em;
  text-align: right;
`;

const Hit = styled.div`
  display: flex;
  align-items: center;
  margin: 5px 30px 5px;
`;

const HitIcon = styled.div`
  width: 55px;
  height: 55px;
  border-radius: 10px;
  background: ${({ theme }) => theme.layoutLittleLessDark};
  box-shadow: 0px 0px 12.1858px rgba(0, 0, 0, 0.25);
`;

const ListItemInfo = styled.div`
  display: flex;
  flex-direction: column;
  margin-left: 10px;
  overflow: hidden;
  text-overflow: ellipsis;
`;

const HitCountSpan = styled.span`
  display: inline-block;
`;

const HitSpan = styled.span`
  display: inline-block;
  color: ${({ theme }) => theme.textDarkest};
`;

const Link = styled.a`
  margin: 3px 0;
  text-decoration: none;
  color: ${({ theme }) => theme.textDarkest};
  transition: color 0.2s ease-in-out;

  &:hover {
    color: ${({ theme }) => theme.textLight};
  }
`;

export const Leaderboard = () => {
  const { hits, isError, isLoading } = useLeaderboard();

  return (
    <Container>
      <Title>Leaderboard</Title>
      <Subtitle>Top 10</Subtitle>
      <br />
      {isLoading && <span>Loading hits</span>}
      {isError && <span>Error loading hits!</span>}

      {hits && hits.length > 0 ? (
        <List>
          {hits.map((hit, i) => {
            return (
              <ListItem key={i}>
                <Place>{i + 1}</Place>
                <Hit>
                  <HitIcon></HitIcon>
                  <ListItemInfo>
                    <Link href={`https://${hit.url}/`}>https://{hit.url}/</Link>
                    <HitCountSpan>
                      {hit.hits} <HitSpan>hits</HitSpan>
                    </HitCountSpan>
                  </ListItemInfo>
                </Hit>
              </ListItem>
            );
          })}
        </List>
      ) : (
        <p>There are no hits yet!</p>
      )}
    </Container>
  );
};
