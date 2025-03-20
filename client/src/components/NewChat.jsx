import React, { useState } from 'react'
import Button from 'react-bootstrap/Button';
import Card from 'react-bootstrap/Card';
import Modal from 'react-bootstrap/Modal';
import InputGroup from 'react-bootstrap/InputGroup';
import Form from 'react-bootstrap/Form';
import 'bootstrap/dist/css/bootstrap.min.css';
import { useForm } from "react-hook-form";
import { useNavigate } from 'react-router-dom';
import { connectWebSocket } from "../utils/websocket";

export const NewChat = () => {
  const navigate = useNavigate();

  const { register, handleSubmit, setError, reset, formState: { errors, isSubmitting, isSubmitted } } = useForm();

  const [show, setShow] = useState(false);

  const handleShow = () => setShow(true);
  const handleClose = () => {
    setShow(false);
    reset();
  }

  const onSubmit = async (data) => {  
    
    console.log("Here 1: ", data);
    let req = await fetch('http://localhost:8070/api/v1/chats', {method: 'POST', headers: {'Content-Type': 'application/json'}, body: JSON.stringify(data)});
    let res = await req.text();

    console.log(res);
    res = JSON.parse(res);
    if (!res.success) {
      setError("myform", {message: res.error});
      return;
    }

    console.log("Here 2: ", res);
    navigate(`/chats/${res.data.chat_room_id}`, { state: {chat_id: res.data.chat_room_id, chat_name: data.chat_room_name, username: data.admin_name} });
    console.log("Here 3: ", res);

    // Websocket connection
    await connectWebSocket(res.data.chat_room_id, data.admin_name);
    let req_ws = await fetch($`http://localhost:8071/ws?chat_id=${res.data.chat_room_id}`, {method: 'POST', headers: {'Content-Type': 'application/json'}, body: JSON.stringify(data)});
    let res_ws = await req.text();    
  }

  return (
    <>
      <Card style={{ width: '18rem' }}>
        <Card.Body>
          <Card.Title>Create Chat Room</Card.Title>
          <Card.Text>
            Give a home to your ideas. Invite friends & hold discussions.
          </Card.Text>
          <Button variant="primary" onClick={handleShow}>Create Room</Button>
        </Card.Body>
      </Card>

      <Modal show={show}>
        <Modal.Header>
          <Modal.Title>Create a chat room</Modal.Title>
        </Modal.Header>
        
        <Modal.Body>
          <Form>
            <InputGroup >
              <Form.Control 
                {...register("admin_name", { required: {value: true, message: "This field is required"}, minLength: {value: 3, message: "Minimum length is 3 characters"}, maxLength: {value: 20, message: "Maximum length is 20 characters"}})}
                className='m-3 bg-light'
                placeholder="Enter a nickname to introduce yourself"
                aria-label="admin_name"
                aria-describedby="basic-addon1"
              />
              {errors.admin_name && (
                  <Form.Text className="text-danger w-100" style={{ fontSize: '0.800rem', marginTop: '-15px', marginLeft: '20px'}}>
                    {errors.admin_name.message}
                  </Form.Text>
              )}
            </InputGroup>
            
            <InputGroup >
              <Form.Control
                {...register("chat_room_name", { required: {value: true, message: "This field is required"}, minLength: {value: 3, message: "Minimum length is 3 characters"}, maxLength: {value: 20, message: "Maximum length is 20 characters"}})}
                className='m-3 bg-light'
                placeholder="Enter a name for your room"
                aria-label="chat_room_name"
                aria-describedby="basic-addon1"
              />
              {errors.chat_room_name && (
                <Form.Text className="text-danger w-100" style={{ fontSize: '0.800rem', marginTop: '-15px', marginLeft: '20px'}}>
                  {errors.chat_room_name.message}
                </Form.Text>
              )}
            </InputGroup>
          </Form>

          {errors.myform && (
              <div className='text-danger'>
                {errors.myform.message}
              </div>
            )}
        </Modal.Body>
        
        <Modal.Footer>
        <Button disabled={isSubmitting} variant="secondary" onClick={handleClose}>
            Cancel
        </Button>
        <Button disabled={isSubmitting} variant="primary" onClick={handleSubmit(onSubmit)}>
            Continue
        </Button>
        </Modal.Footer>
      </Modal>
    </>
  )
}

export default NewChat