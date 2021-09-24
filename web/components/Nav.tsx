import styled from "styled-components";

const Container = styled.div`
  width: 100%;
  padding: 10px 20px;
  display: flex;
  align-items: center;
`;

const Brand = styled.h1`
  margin-right: 15px;
  font-weight: black;
`;

const BrandSpan = styled.span`
  font-size: 1.2em;
  font-weight: medium;
  color: ${({ theme }) => theme.textDarkest};
  flex-grow: 1;
`;

const Creators = styled.div`
  display: flex;
  flex-direction: column;
  align-items: flex-end;
`;

const Links = styled.a`
  margin: 3px 0;
  text-decoration: none;
  color: ${({ theme }) => theme.textDarkest};
  transition: color 0.2s ease-in-out;

  &:hover {
    color: ${({ theme }) => theme.textLight};
  }
`;

export const Nav = () => {
  return (
    <Container>
      <Brand>Hits</Brand>
      <BrandSpan>The better hit counter</BrandSpan>
      <Creators>
        <Links
          href="https://twitter.com/heybereket"
          target="_blank"
          rel="noreferrer"
        >
          @heybereket
        </Links>
        <Links
          href="https://twitter.com/devlooskie"
          target="_blank"
          rel="noreferrer"
        >
          @devlooskie
        </Links>
      </Creators>
    </Container>
  );
};
