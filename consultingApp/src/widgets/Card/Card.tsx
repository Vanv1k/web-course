import Button from 'react-bootstrap/Button';
import CardBootstrap from 'react-bootstrap/Card';
import { Link } from 'react-router-dom';

const Card = (props) => {

    return (
        <CardBootstrap style={{ width: '18rem', marginTop: '3rem', margin:'10% 10% 5% 10%' }}>
            <CardBootstrap.Img variant="top" src={props.image} />
            <CardBootstrap.Body>
                <CardBootstrap.Title>{props.name} </CardBootstrap.Title>
                <CardBootstrap.Text>
                    {props.description}
                </CardBootstrap.Text>
                {/* <Link to={`/consultations/${props.id}`}> */}
                    <Button variant="primary">Go somewhere</Button>
                {/* </Link> */}

                <CardBootstrap.Text>
                    {props.price} рублей
                </CardBootstrap.Text>
            </CardBootstrap.Body>
        </CardBootstrap>
    );
}

export default Card;