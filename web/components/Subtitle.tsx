import React from "react";
import styled from "styled-components";

export const H1 = styled.h1`
  margin: 0;
  padding: 0;
  font-size: 1.15em;
  font-weight: 600;
  color: ${({ theme }) => theme.textDarkest};
`;

export const Subtitle = ({ children }: { children: JSX.Element | string }) => {
  return <H1>{children}</H1>;
};
