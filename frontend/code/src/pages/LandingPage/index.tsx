import React, { useState } from "react";

import {
  Container,
  LeftContainer,
  RightContainer,
  ButtonsContainer,
  AuthButton,
} from "./styles";
import MonitoringImage from "../../assets/monitoring.png";
import RegisteredImage from "../../assets/registered.png";

import { LoginForm } from '../Auth/Login';
import { RegisterForm } from '../Auth/Register';

export const LandingPage = () => {
  const [activeStep, setActiveStep] = useState(0);

  const backToHomeScreen = (): void => {
    setActiveStep(0)
  }

  const renderActiveStep = () => {
    let actualRender;
    switch(activeStep) {
      case 1:
        actualRender = <LoginForm backToHomeScreen={backToHomeScreen} />
        break
      case 2:
          actualRender = <RegisterForm backToHomeScreen={backToHomeScreen} />
          break
      default:
        actualRender = (<>
          <div id="message">
            <p>Faça login para usar comar a analisar seus dados</p>
          </div>
          <ButtonsContainer>
            <AuthButton onClick={() => setActiveStep(1)}>Login</AuthButton>
            <AuthButton onClick={() => setActiveStep(2)}>Registrar</AuthButton>
          </ButtonsContainer>
        </>)
    }
    return actualRender;
  }

  return (
    <Container>
      <LeftContainer>
        <h1>Análise de espectros</h1>
        <img src={MonitoringImage} alt="logo" />
        <div>
          <p>
            Analise os espectros de suas amostras de forma automatizada
            utilizando aprendizado de máquina
          </p>
        </div>
      </LeftContainer>
      <RightContainer>
        <img src={RegisteredImage} alt="login" />
        {renderActiveStep()}
      </RightContainer>
    </Container>
  );
};
