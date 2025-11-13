function Card({num}) {
    return (
        <div className="card">
            <h1>This card's value is {num}</h1>
            <p>The card is :{ num>50?"High":"Low"}</p>
        </div>
    )
}
export default Card;