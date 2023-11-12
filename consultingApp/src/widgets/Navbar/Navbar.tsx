import './Navbar.css'
import { ChangeEvent } from 'react';
import Button from 'react-bootstrap/Button';
import Container from 'react-bootstrap/Container';
import Form from 'react-bootstrap/Form';
import Nav from 'react-bootstrap/Nav';
import { Navbar as NavB } from 'react-bootstrap';
import { useState } from 'react';

interface NavbarProps {
  onMaxPriceChange?: (value: string) => void; // Define the prop type
}

const Navbar: React.FC<NavbarProps> = ({ onMaxPriceChange }) => {
  const [maxPrice, setMaxPrice] = useState('');

  const handleMaxPriceChange = (e: ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    setMaxPrice(value);

    // Check if onMaxPriceChange is defined before calling it
    if (onMaxPriceChange !== undefined) {
      onMaxPriceChange(value);
    }
  };

  return (
    <NavB expand="lg" bg="dark" data-bs-theme="dark" className="bg-body-tertiary">
      <Container fluid style={{ marginLeft: '5%' }}>
        <NavB.Brand href="/">IT Services</NavB.Brand>
        <NavB.Toggle aria-controls="navbarScroll" />
        <NavB.Collapse id="navbarScroll">
          <Nav
            className="me-auto my-2 my-lg-0"
            style={{ maxHeight: '100px' }}
            navbarScroll
          >
            <Nav.Link href="/">Главная</Nav.Link>
            <Nav.Link href="#action2">Корзина</Nav.Link>
          </Nav>
          <Form
            className="d-flex"
            id="search"
          >
            <Form.Control
              type="search"
              placeholder="Поиск по максимальной цене"
              className="me-2"
              aria-label="Search"
              value={maxPrice}
              onChange={handleMaxPriceChange}
            />
            <Button variant="outline-success" type="submit">
              Искать
            </Button>
          </Form>
        </NavB.Collapse>
      </Container>
    </NavB>
  );
};

export default Navbar;
