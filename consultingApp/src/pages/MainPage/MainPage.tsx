import './MainPage.css'
import Navbar from '../../widgets/Navbar/Navbar'
import Card from '../../widgets/Card/Card'



const MainPage = () => {

    const data = [
        {
            id: 1,
            Name: 'ffsrfa',
            Description: 'aergergaerg',
            Image: 'https://placekitten.com/300/200',
            Price: 12312321,
            Status: 'active',
        },
        {
            id: 2,
            Name: 'f13fsrfa',
            Description: 'aergergaerg',
            Image: 'https://placekitten.com/300/200',
            Price: 12312321,
            Status: 'active',
        },
        {
            id: 3,
            Name: 'f152fsrfa',
            Description: 'aergergaerg',
            Image: 'https://placekitten.com/300/200',
            Price: 12312321,
            Status: 'active',
        },
        {
            id: 4,
            Name: 'ffa',
            Description: 'aergergaerg',
            Image: 'https://placekitten.com/300/200',
            Price: 12312321,
            Status: 'active',
        },
        {
            id: 5,
            Name: 'f13fsrfa',
            Description: 'aergergaerg',
            Image: 'https://placekitten.com/300/200',
            Price: 12312321,
            Status: 'active',
        },
        {
            id: 6,
            Name: 'f152fsrfa',
            Description: 'aergergaerg',
            Image: 'https://placekitten.com/300/200',
            Price: 12312321,
            Status: 'active',
        },
    ]

    return (
        <div>
            <Navbar />
            <div className="container">
                <div className="row">
                    {data.map((item) => (
                        <div key={item.id} className="col-lg-3 col-md-4 col-sm-6">
                            <Card
                                id={item.id}
                                name={item.Name}
                                description={item.Description}
                                image={item.Image}
                                price={item.Price}
                            />
                        </div>
                    ))}
                </div>
            </div>
        </div>
    )

}

export default MainPage 