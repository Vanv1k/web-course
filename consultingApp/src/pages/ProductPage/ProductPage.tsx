import './ProductPage.css';
import Navbar from '../../widgets/Navbar/Navbar';
import Button from 'react-bootstrap/Button';
import { useParams } from 'react-router-dom';
import { Container, Row, Col } from 'react-bootstrap';

const ProductPage = (props) => {
    const data = {
        id: 1,
        Name: 'name',
        Description: 'desc',
        Image: 'https://placekitten.com/300/200',
        Price: 12312321,
        Status: 'active',
    }

    const { cardId } = useParams(); // Получаем параметры из URL

    return (
        <div>
            <Navbar />
            <Container style={{marginTop: '10%'}}>
                <Row>
                    <Col xs={12} md={6}>
                        <img src={data.Image} className="card-img-selected" alt={data.Name}  style={{ borderRadius: '10px' }} />
                    </Col>
                    <Col xs={12} md={6}>
                    <h1 className="text card-name-selected" style={{fontSize: '150%',fontWeight: 'bold' }}>{data.Name}</h1>
                        <p className="text card-description-selected">{data.Description}</p>
                        <div className="bottom-part">
                            <p className="text card-price-selected">{data.Price} рублей</p>
                            <Button variant="primary">Провести</Button>
                        </div>
                    </Col>
                </Row>
            </Container>
        </div>
    )
}

export default ProductPage;