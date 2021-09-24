import useSWR from "swr";
import { Hit } from "../types";
import { FULL_URI_V1 } from "../utils";

export const useLeaderboard = () => {
  const { data, error } = useSWR(`${FULL_URI_V1}/top`);
  const hits: Array<Hit> = data?.data;

  return {
    hits: hits,
    isLoading: !error && !data,
    isError: error,
  };
};
