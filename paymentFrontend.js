import Request from ".././api/requests";
import { useNavigate } from "react-router-dom";
import { useEffect } from "react";

const PayByRazor = (props) => {
    const navigate = useNavigate();

    const openPaymentModal = async () => {
        props.handleLoading(true);
        let { data } = await Request.createOrder({ eventId: props?.eventDetails?._id });

        if (props?.eventDetails?.fees === 0){
            navigate("/events/" + props?.eventDetails?._id)
            return
        }

        const options = {
            key: process.env.REACT_APP_RAZORPAY_KEY,
            amount: data?.amount.toString(),
            currency: "INR",
            name: "PCSB Xenia 2022",
            description: "Registration payment for the event " + props.eventDetails?.eventName || "",
            order_id: data?.id,
            callback_url: "",
            prefill: {
                name: props.userDetails?.name || "",
                mobile: props.userDetails?.mobile.toString() || ""
            },
            notes: {
                eventId: props?.eventDetails?._id || "",
                userId: data?.userId || "",
                teamName: props?.teamName || "",
            },
            theme: {
                "color": "#00a0fd83"
            }
        }

        const razorDisplay = new window.Razorpay(options);
        props.handleLoading(false)
        razorDisplay.open();
    }

    useEffect(() => {
        const script = document.createElement('script');
        script.src = 'https://checkout.razorpay.com/v1/checkout.js';
        script.async = true;
        document.body.appendChild(script);
    }, [])

    return (
        <>
            <button className={props.className} onClick={openPaymentModal} type="submit">{props.buttonName}</button>
        </>
    )
}

export default PayByRazor;