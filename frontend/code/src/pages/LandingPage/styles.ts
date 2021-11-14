import styled from 'styled-components';


export const Container = styled.div`
  height: 100%;
  width: 100%;
  display: flex;
  flex-direction: row;
  background-color: #f6f7f2;
`;

export const LeftContainer = styled.div`
  width: 60%;
  background-color: #fafafa;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;

  > img {
    width: 500px;
    border-radius: 80px;
  }

  > h1 {
    margin-bottom: 20px;
    font-size: 50px;
    color: #222;
  }

  > div {
    width: 70%;
    display: flex;
    align-items: center;
    justify-content: center;

    > p {
      text-align: center;
      margin-top: 30px;
      font-size: 20px;
      color: #333;
    }
  }
`;

export const RightContainer = styled.div`
  width: 40%;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;

  > img {
    width: 120px;
  }
  > div#message {
    margin-top: 35px;
    width: 80%;
    display: flex;
    align-items: center;
    justify-content: center;
    
    > p {
      font-size: 20px;
      color: #333;
    }
  }
`;

export const ButtonsContainer = styled.div`
  width: 100%;
  margin-top: 20px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
`;

export const AuthButton = styled.button`
  width: 50%;
  height: 50px;
  color: white;
  border: 0;
  margin: 5px 0;
  border-radius: 5px;
  font-size: 18px;
  cursor: pointer;
  background-color: #548AFF;
`