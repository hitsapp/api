import useSWR from "swr";
import { Hit } from "../types";
import { API_URL } from "../utils";

export const useLeaderboard = () => {
  const { data, error } = useSWR(`${API_URL}/top`);
  const hits: Array<Hit> = data?.data;

  return {
    hits: hits,
    isLoading: !error && !data,
    isError: error,
  };
};
