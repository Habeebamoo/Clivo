import { Outlet } from "react-router"
import Header from "../components/Header"

const Layout = () => {
  return (
    <>
      <Header type="dashboard" />
      <Outlet />
    </>
  )
}

export default Layout