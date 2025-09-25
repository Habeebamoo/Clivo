import { H2 } from "../../components/Typo"
import logo from "../../assets/logo.jpg"
import googleImg from "../../assets/google.png"
import Button from "../../components/Button"

const LoginPage = () => {
  return (
    <main className="flex-center flex-col h-[100vh]">
      <img src={logo} className="h-15 rounded-md mb-6" />
      <H2 font="inter" text="Sign In With Google" />
      <Button text="Continue With Google" img={googleImg} imgClass="h-5" />
    </main>
  )
}

export default LoginPage