import Button from 'react-bootstrap/Button';
import CardBootstrap from 'react-bootstrap/Card';
import { Link } from 'react-router-dom';

interface CardProps {
    id: number;
    name: string;
    description: string;
    image: string;
    price: number;
}

const MAX_DESCRIPTION_LENGTH = 100; 

const Card: React.FC<CardProps> = (props) => {
    const truncatedDescription = props.description.length > MAX_DESCRIPTION_LENGTH
        ? `${props.description.substring(0, MAX_DESCRIPTION_LENGTH)}...`
        : props.description;

    return (
        <CardBootstrap style={{ width: '18rem', marginTop: '3rem', margin: '10% 10% 5% 10%' }}>
            <CardBootstrap.Img variant="top" src={props.image} />
            <CardBootstrap.Body>
                <CardBootstrap.Title>{props.name} </CardBootstrap.Title>
                <CardBootstrap.Text>
                    {truncatedDescription}
                </CardBootstrap.Text>
                <>
                    <Link to={`/consultations/${props.id}`} style={{ marginRight: '10px' }}>
                        <Button variant="primary">Подробнее</Button>
                    </Link>
                    <CardBootstrap.Text style={{ display: 'inline-block', fontWeight: 'bold', marginTop: '2%' }}>
                        {props.price} рублей
                    </CardBootstrap.Text>
                </>
            </CardBootstrap.Body>
        </CardBootstrap>
    );
}

export default Card;