const data={
    heading:"99% of all times",
    callToAction:"Everything must go"
}

function Promo(){
    return(
        <>
        <PromoHandling heading={data.heading} callToAction={data.callToAction}/>
        </>
    )
}

export default Promo