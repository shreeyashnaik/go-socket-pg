import React, { useState } from 'react'
import Button from 'react-bootstrap/Button';
import Card from 'react-bootstrap/Card';
import Modal from 'react-bootstrap/Modal';
import 'bootstrap/dist/css/bootstrap.min.css';
import InputGroup from 'react-bootstrap/InputGroup';
import Form from 'react-bootstrap/Form';
import { useForm } from "react-hook-form";

const JoinChat = () => {
    const { register, handleSubmit, setError, reset, formState: { errors, isSubmitting, isSubmitted } } = useForm();

    const [show, setShow] = useState(false);
  
    const handleShow = () => setShow(true);
    const handleClose = () => {
      setShow(false);
      reset();
    }
  
    const onSubmit = async (data) => {  
  
      try {
        let req = await fetch(`http://localhost:8070/api/v1/chats/${data.chat_room_id}/users`, {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ username: data.username })
        });

        let res = await req.json();

        if (!res.status) {
          setError("myform", { message: res.error });
          return;
        }

        setShow(false);

        console.log(res);

        navigate(`/chats/${data.chat_room_id}`);
      } catch (error) {
        console.error("Error:", error);
        setError("myform", { message: "An unexpected error occurred. Please try again later." });
      }
    }
  
    return (
      <>
        <Card style={{ width: '18rem' }}>
          <Card.Body>
            <Card.Title>Join Chat Room</Card.Title>
            <Card.Text>
                Join a chat room created by your friend. Get a link or code from them.
            </Card.Text>
            <Button variant="primary" onClick={handleShow}>Join Room</Button>
          </Card.Body>
        </Card>
  
        <Modal show={show}>
          <Modal.Header>
            <Modal.Title>Join Chat Room</Modal.Title>
          </Modal.Header>
          
          <Modal.Body>
            <Form>
              <InputGroup >
                <Form.Control 
                  {...register("chat_room_id", { required: {value: true, message: "This field is required"}})}
                  className='m-3 bg-light'
                  placeholder="Enter a code or link"
                  aria-label="chat_room_id"
                  aria-describedby="basic-addon1"
                />
                {errors.chat_room_id && (
                    <Form.Text className="text-danger w-100" style={{ fontSize: '0.800rem', marginTop: '-15px', marginLeft: '20px'}}>
                      {errors.chat_room_id.message}
                    </Form.Text>
                )}
              </InputGroup>
              
              <InputGroup >
                <Form.Control
                  {...register("username", { required: {value: true, message: "This field is required"}, minLength: {value: 3, message: "Minimum length is 3 characters"}, maxLength: {value: 20, message: "Maximum length is 20 characters"}})}
                  className='m-3 bg-light'
                  placeholder="Enter a nickname to introduce yourself"
                  aria-label="username"
                  aria-describedby="basic-addon1"
                />
                {errors.username && (
                  <Form.Text className="text-danger w-100" style={{ fontSize: '0.800rem', marginTop: '-15px', marginLeft: '20px'}}>
                    {errors.username.message}
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

export default JoinChat;