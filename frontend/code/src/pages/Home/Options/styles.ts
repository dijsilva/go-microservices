import styled from "styled-components";

export const ContainerOptions = styled.div`
  width: 100%;
  height: 92%;
  background-color: #fafafa;
  display: flex;
  align-items: center;
  justify-content: center;
`;


export const ContainerOption = styled.div`
  width: 300px;
  height: 300px;
  background-color: #fafafa;
  display: flex;
  cursor: pointer;
  align-items: center;
  flex-direction: column;
  justify-content: center;
  border-radius: 10px;
  border: 1px solid #999;
  justify-content: center;
  margin: 0px 25px;

  > p {
    margin-top: 20px;
    font-size: 20px;
    font-weight: bold;
    color: #333
  }
`;