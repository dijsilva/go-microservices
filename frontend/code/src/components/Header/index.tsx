import React, { useEffect, useState } from "react";
import { useNavigate } from 'react-router-dom';

import { HeaderContainer } from "./styles";
import { FiLogOut } from 'react-icons/fi';
import { HiHome } from 'react-icons/hi';

export const Header = (): JSX.Element => {
  const navigate = useNavigate()
  const [userData, setUserData] = useState({
    name: "",
    userName: "",
  });

  useEffect(() => {
    const user = localStorage.getItem("user");

    if (user) {
      setUserData({ ...JSON.parse(user) });
    }
  }, []);

  const logout = () => {
    localStorage.removeItem('user');
    localStorage.removeItem('token');
    navigate('/')
  }

  return (
    <HeaderContainer>
        <p>Olá, <strong>{userData.name}</strong></p>
        <div>
          <HiHome onClick={() => navigate('/home')} size={28} title="Sair" color="white" />
          <FiLogOut onClick={() => logout()} size={28} title="Sair" color="white" />
        </div>
    </HeaderContainer>
  );
};
