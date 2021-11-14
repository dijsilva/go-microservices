import React, { useState } from "react";

import { Header } from '../../components/Header';
import {
  Container,
  Form,
  TextInput,
  LabelInput,
  DropZoneArea,
  DropZoneContainer
} from './styles';

import { useNavigate } from 'react-router-dom';

import { useDropzone } from 'react-dropzone'
import { GiConfirmed } from 'react-icons/gi';

import spectraService from '../../services/spectra';
import axios from "axios";
import { toast } from "react-toastify";

export const NewSpectra = (): JSX.Element => {
  const navigate = useNavigate()
  const [waitingToSend, setWaitingToSend] = useState(true);
  const [spectraFile, setSpectraFile] = useState<File>()
  const [spectraFilename, setSpectraFileName] = useState<string>('')
  const [valuesToCreateSpectra, setValuesToCreateSpectra] = useState({
    sample_name: '',
    n_spectra: 0,
    equipament_used: '',
  })

  function StyledDropzone(props: any) {
    const {
      getRootProps,
      getInputProps,
      isDragActive,
      isDragAccept,
      acceptedFiles,
      isDragReject
    } = useDropzone({ maxFiles: 1, accept: ".csv" });

    if (acceptedFiles && acceptedFiles.length > 0) {
      setSpectraFile(acceptedFiles[0])
      setSpectraFileName(acceptedFiles[0].name)
    }
    
    return (
      <DropZoneContainer>
        <DropZoneArea {...getRootProps({isDragActive, isDragAccept, isDragReject})}>
          <input {...getInputProps()} />
          {<p>{spectraFilename ? `${spectraFilename}` : `Arraste e solte aqui o seu arquivo .csv com os espectros`}</p>}
        </DropZoneArea>
      </DropZoneContainer>
    );
  }

  const validFormFields = (input: typeof valuesToCreateSpectra): boolean => {
    let isValid = true;
    Object.entries(input).forEach(([_, value]) => {
      if (!value) {
        isValid = false;
      }
    })
    return isValid;

  }

  const updateFormValues = (input: React.ChangeEvent<HTMLInputElement>): void => {
    const { value, name } = input.target;

    setValuesToCreateSpectra({...valuesToCreateSpectra, [name]: value})
  }

  const onSubmitForm = async (): Promise<void> => {
    if (!validFormFields(valuesToCreateSpectra)) {
      toast.error('Os campos dos formulários não podem ficar vazios.')
      return
    }
    const formData = new FormData()
    formData.append('sample_name', valuesToCreateSpectra.sample_name)
    formData.append('equipment_used', valuesToCreateSpectra.equipament_used)
    formData.append('n_spectra', String(valuesToCreateSpectra.n_spectra))
    if (spectraFile) {
      formData.append('spectra_file', spectraFile)
    }

    let token;
    token = localStorage.getItem('token')
    if (token) {
      token = JSON.parse(token)
    }

    let response;
    try {
      response = await spectraService.post('/create', formData, {
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'multipart/form-data'
        }
      })
      if (response && response.status && response.status === 201) {
        toast.success('Espectro enviado!')
        setTimeout(() => {
          navigate('/home')
        }, 3000)
      }
    } catch (err) {
      if (axios.isAxiosError(err)) {
        console.log(err.response?.data)
      }
    }
  }
  return (
    <>
    <Header />
    <Container>
      <div>
        <Form onSubmit={onSubmitForm}>
        <div>
          <LabelInput>Amostra</LabelInput>
          <TextInput
            type="text"
            name="sample_name"
            maxLength={150}
            placeholder="Digite uma descrição da amostra"
            onChange={updateFormValues}
            value={valuesToCreateSpectra.sample_name}
          />
        </div>
        <div>
          <LabelInput>Número de espectros</LabelInput>
          <TextInput
            type="number"
            step={1}
            min={1}
            name="n_spectra"
            placeholder="Insira o número de espectros na amostra"
            onChange={updateFormValues}
            value={valuesToCreateSpectra.n_spectra}
          />
        </div>
        <div>
          <LabelInput>Equipamento</LabelInput>
          <TextInput
            type="email"
            name="equipament_used"
            placeholder="Digite seu equipamento em que as amostras foram lidas"
            onChange={updateFormValues}
            value={valuesToCreateSpectra.equipament_used}
          />
        </div>
        </Form>
        <StyledDropzone />

      </div>
      <div>
        {waitingToSend ? <GiConfirmed size={120} color="#52cf15" onClick={async () => await onSubmitForm()} title="Enviar espectro" /> : <></>}
      </div>
    </Container>
    </>
  );
};
