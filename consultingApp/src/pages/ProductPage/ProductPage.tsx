import './ProductPage.css';
import Navbar from '../../widgets/Navbar/Navbar';
import Button from 'react-bootstrap/Button';
import { useParams } from 'react-router-dom';
import { Container, Row, Col } from 'react-bootstrap';
import { useState, useEffect } from 'react';

interface ProductData {
    id: number;
    Name: string;
    Description: string;
    Image: string;
    Price: number;
    Status: string;
  }

  const ProductPage: React.FC = () => {
    const { id } = useParams();
    console.log(id)

    const [data, setData] = useState<ProductData | null>(null);

    useEffect(() => {
      // Выполняем запрос при монтировании компонента
      fetchData();
    }, []);
  
    const fetchData = async () => {
      try {
        const response = await fetch(`/api/consultations/${id}`);
        if (!response.ok) {
          throw new Error(`Ошибка при выполнении запроса: ${response.statusText}`);
        }
  
        const result = await response.json();
        setData(result);
      } catch (error) {
        console.error('ошибка при выполннении запроса:', error);
      }
    };
    console.log(data);

    return (
        <div>
            <Navbar />
            <Container style={{marginTop: '10%'}}>
                <Row>
                    <Col xs={12} md={6}>
                        <img src={data?.Image} className="card-img-selected" alt={data?.Name}  style={{ borderRadius: '10px' }} />
                    </Col>
                    <Col xs={12} md={6}>
                    <h1 className="text card-name-selected" style={{fontSize: '150%',fontWeight: 'bold' }}>{data?.Name}</h1>
                        <p className="text card-description-selected">{data?.Description}</p>
                        <div className="bottom-part">
                            <p className="text card-price-selected">{data?.Price} рублей</p>
                            <Button variant="primary">Провести</Button>
                        </div>
                    </Col>
                </Row>
            </Container>
        </div>
    )
}

export default ProductPage;