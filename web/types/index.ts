export interface Hits {
  [index: number]: Hit;
}
export interface Hit {
  url: string;
  hits: number;
}

export interface Response {
  success: boolean;
  message?: string;
  data?: object;
}
