import styled from "styled-components";

export const ListRequestsContainer = styled.div`
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  height: 92%;
  width: 100%;
  padding: 10px;
`;

export const RequestContainer = styled.div`
  width: 320px;
  height: 350px;
  border: 1px solid #444;
  display: flex;
  margin: 10px;
  align-items: flex-start;
  justify-content: space-between;
  flex-direction: column;
  padding: 8px;
  border-radius: 4px;

  div#title {
    padding: 5px;
    border-bottom: 1px solid #999;
    margin-bottom: 10px;
  }

  h4 {
    font-weight: 400;
  }

  span {
    text-transform: uppercase;
  }
`;
