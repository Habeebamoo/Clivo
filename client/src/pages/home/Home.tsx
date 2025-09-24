import Header from "../../components/Header"
import PhotoGrid from "../../components/PhotoGrid"
import { H1 } from "../../components/Typo"

const Home = () => {
  return (
    <>
      <Header />

      {/* Hero section */}
      <div className="mt-18 pt-10 px-4">
        <H1 font="inter" text="Where Simple Stories Find Their Voices" others="md:text-center" />
        <p className="text-accent mt-2 mb-10 text-sm md:mx-auto md:text-center md:w-[70%]">Clivo is a platform for writers who demands excellence. Craft, Publish, and share your stories with a community that values quality</p>
        <PhotoGrid />
      </div>

      {/* Subscription Section */}
      <div className="mt-20">

      </div>
    </>
  )
}

export default Home