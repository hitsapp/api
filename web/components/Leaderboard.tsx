import { useLeaderboard } from "../hooks/useLeaderboard";
import styled from "styled-components";
import { Subtitle, Title } from ".";
import Skeleton from "react-loading-skeleton";

const Container = styled.div`
  padding: 24px 24px;
  width: 410px;
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

  const messages = [
    "Fetching top hits...",
    `Who's in top 10?`,
    "bereket was here",
    "looskie was here",
    "Loading results...",
  ]
  
  const generateMessage = (array: string[]) => {
    const message = Math.floor(Math.random() * array.length);
    return array[message];
  }

  return (
    <Container>
      <Title>Leaderboard</Title>
      <Subtitle>Top 10</Subtitle>
      <br />
      {isLoading || isError && (
        <List>
          <p style={{ color: "#9E9E9E" }}>{generateMessage(messages)}</p>
          {[...new Array(20)].map((_, i) => {
            return (
              <ListItem style={{ margin: "8px 0px" }} key={i + 1}>
                <Place>{i + 1}</Place>
                <Skeleton
                  style={{ marginLeft: 25, borderRadius: 10 }}
                  width={55}
                  height={55}
                />
                <ListItemInfo>
                  <Skeleton width={140} height={15} />
                  <HitCountSpan>
                    <Skeleton width={30} height={15} />
                    <Skeleton
                      style={{ marginLeft: 10 }}
                      width={30}
                      height={15}
                    />
                  </HitCountSpan>
                </ListItemInfo>
              </ListItem>
            );
          })}
        </List>
      )}

      {hits && hits.length > 0 ? (
        <List>
          {hits.map((hit, i) => {
            return (
              <ListItem key={i}>
                <Place>{i + 1}</Place>
                <Hit>
                  <HitIcon></HitIcon>
                  <ListItemInfo>
                    <Link href={hit.url} target="_blank">
                      {hit.url}
                    </Link>
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
        console.log("[HITS] - No hits recieved!")
      )}
    </Container>
  );
};
