import React from "react";
import styled from "styled-components";

export const H1 = styled.h1`
  margin: 0 0 0.5rem 0;
  padding: 0;
`;

export const Title = ({ children }: { children: JSX.Element | string }) => {
  return <H1>{children}</H1>;
};
