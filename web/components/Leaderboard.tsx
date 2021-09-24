import { useLeaderboard } from "../hooks/useLeaderboard";

export const Leaderboard = (): JSX.Element => {
  const { hits, isError, isLoading } = useLeaderboard();

  console.log(isError);
  return (
    <div>
      <h1>Leaderboard</h1>
      {isLoading && <span>Loading hits</span>}
      {isError && <span>Error loading hits!</span>}

      {hits && hits.length > 0 ? (
        <ul>
          {hits.map((hit, i) => {
            return (
              <li key={i}>
                {hit.hits}
                <br />
                {hit.url}
              </li>
            );
          })}
        </ul>
      ) : (
        <p>There are no hits yet!</p>
      )}
    </div>
  );
};
