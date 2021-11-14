import styled from "styled-components";

export const TextInput = styled.input`
  height: 50px;
  padding: 0px 10px;
  width: 100%;
  border: 0.6px solid #999;
  border-radius: 5px;
  background-color: #e7e8e1;
  font-size: 20px;
`;

export const LabelInput = styled.p`
  font-size: 22px;
  margin-left: 2px;
  width: 100%;
  color: #333;
`;

export const Form = styled.form`
  width: 100%;
  display: flex;
  margin-top: 40px;
  align-items: center;
  justify-content: center;
  flex-direction: column;

  > div {
    width: 70%;
    margin-bottom: 20px;
    display: flex;
    align-items: center;
    justify-content: left;
    flex-direction: column;
  }
`;

export const RegisterButton = styled.button`
  display: flex;
  align-items: center;
  justify-content: center;
  width: 50%;
  height: 50px;
  color: white;
  border: 0;
  margin: 5px 0;
  border-radius: 5px;
  font-size: 20px;
  font-weight: bold;
  cursor: pointer;
  background-color: #548AFF;
`