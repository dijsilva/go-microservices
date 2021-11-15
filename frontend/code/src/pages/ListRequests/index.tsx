import React, { useEffect, useState } from "react";

import { Header } from "../../components/Header";

import { ListRequestsContainer, RequestContainer, StatusText } from './styles';

import spectraService from '../../services/spectra';
import axios from "axios";
import { toast } from "react-toastify";

interface RequestsList {
  id: string;
  sample_name:string;
  email_owner:string;
  n_spectra: number;
  equipament_used: string;
  prediction_concluded: boolean;
  created_at: Date;
  updated_at: Date;
  prediction_info: {
    prediction_date: Date;
    prediction_string: string;
    prediction_number: number;
  }
}

export const ListRequests = (): JSX.Element => {
  const [spectraRequests, setSpectraRequests] = useState<RequestsList[]>([])

  const getRequests = async () => {
    let response;
    let token = localStorage.getItem('token')
    if (token) {
      token = JSON.parse(token)
    }
    try {
      response = await spectraService.get('list-by-owner', {
        headers: {
          'Authorization': `Bearer ${token}`,
        }
      });
      if (response && response.data && Array.isArray(response.data) && response.data.length > 0) {
        let requests: RequestsList[] = response.data;
        setSpectraRequests([...requests])
      } else {
        toast.error('Ocorreu um erro ao buscar suas solicitações')
      }
    } catch (err) {
      if (axios.isAxiosError(err)) {
        if (err.response && err.response.status && err.response.status === 401) {
          toast.error('Acesso negado. Tente fazer login novamente.')
        } else if (err.response && err.response.status && err.response.status === 404) {
          toast.warn('Nenhuma solicitação encontrada.')
        }
      } else {
        toast.error('Ocorreu um erro no servidor.')
      }
    }
  }

  const getStatus = (predictionStatus: boolean): string => {
    return predictionStatus ? 'ANALISADO' : 'PENDENTE'
  }

  const formatData = (date: Date): string => {
    const dateInstance = new Date(date);
    const dateString = dateInstance.toLocaleDateString('pt-br', {
      day: 'numeric',
      month: 'long',
      year: 'numeric',
    })
    return dateString
  }

  const formatHour = (date: Date): string => {
    const dateInstance = new Date(date);
    const dateString = dateInstance.toLocaleDateString('pt-br', {
      hour: 'numeric',
      minute: 'numeric',
    })
    return dateString.substr(11, 15).replace(':', 'h')
  }

  const calAsyncFunctions = async () => {
    await getRequests();
  }

  useEffect(() => {
    calAsyncFunctions()
  }, []);
  return (
    <>
      <Header />
      <ListRequestsContainer>
        {spectraRequests.length > 0 && spectraRequests.map((request, index) => (
          <RequestContainer>
          <div>
            <div id="title">
              <h2>Solicitação nº{index + 1}</h2>
            </div>
            <h4><strong>Solicitação:</strong> {request.sample_name}</h4>
            <h4><strong>Equipamento:</strong> {request.equipament_used}</h4>
            <h4><strong>Enviado em:</strong> {formatData(request.created_at)}</h4>
            <h4><strong>Hora de envio:</strong> {formatHour(request.created_at)}</h4>
            {request.prediction_info?.prediction_string && (
              <h3 id="result"><strong>Resultado:</strong> {request.prediction_info.prediction_string}</h3>
            )}
          </div>
          <span>Status: <StatusText concluded={request.prediction_concluded}>{getStatus(request.prediction_concluded)}</StatusText></span>
        </RequestContainer>
        ))}
        
      </ListRequestsContainer>
    </>
  );
};
