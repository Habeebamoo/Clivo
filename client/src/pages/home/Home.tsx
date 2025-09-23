import Header from "../../components/Header"
import PhotoGrid from "../../components/PhotoGrid"

const Home = () => {
  return (
    <>
      <Header />

      {/* Hero section */}
      <div className="mt-18 bg-green-200 pt-10">
        <h1 className="font-inter">Where Simple Stories Find Thier Voices</h1>
        <PhotoGrid />
      </div>

      {/* Subscription Section */}
      <div>

      </div>
    </>
  )
}

export default Home