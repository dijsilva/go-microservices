import styled from "styled-components";

export const Container = styled.div`
  width: 100%;
  height: 92%;
  display: flex;
  flex-direction: column;

  > div {
    display: flex;
    height: 50%;
    width: 100%;
    align-items: center;
    justify-content: center;

    > svg {
      cursor: pointer;
    }
  }
  /* align-items: center; */
  /* justify-content: center; */
`;

export const Form = styled.form`
  width: 55%;
  height: 100%;
  padding-right: 20px;
  padding: 0;
  margin: 0;
  outline: 0;
  display: flex;
  flex-direction: column;
  padding-left: 10px;
  padding-top: 20px;
  /* justify-content: center; */
`;

export const TextInput = styled.input`
  height: 50px;
  padding: 0px 10px;
  width: 95%;
  margin-bottom: 30px;
  border: 0.6px solid #999;
  border-radius: 5px;
  background-color: #e7e8e1;
  font-size: 20px;
`;

export const LabelInput = styled.p`
  font-size: 22px;
  margin-left: 2px;
  margin-bottom: 5px;
  width: 95%;
  color: #333;
`;


export const DropZoneContainer = styled.div`
  width: 45%;
  height: 100%;
  display: flex;
  cursor: pointer;
  padding-top: 25px;
  /* align-items: center; */
  justify-content: center;
`;

export const DropZoneArea = styled.div`
  /* flex: 1; */
  height: 75%;
  width: 95%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 20px;
  border-width: 2px;
  border-radius: 2px;
  border-color: ${props => getColor(props)};
  border-style: dashed;
  background-color: #fafafa;
  font-size: 30px;
  color: #bdbdbd;
  outline: none;
  transition: border .24s ease-in-out;
`;

const getColor = (props: any) => {
  if (props.isDragAccept) {
      return '#00e676';
  }
  if (props.isDragReject) {
      return '#ff1744';
  }
  if (props.isDragActive) {
      return '#2196f3';
  }
  return '#eeeeee';
}