import { useState } from "react"
import Header from "../../components/Header"
import PhotoGrid from "../../components/PhotoGrid"
import { H1, H2 } from "../../components/Typo"

const Home = () => {
  const [email, setEmail] = useState<string>("")

  const subscribe = (e: React.FormEvent) => {
    e.preventDefault()
    if (!email) return
    alert(email)
  }

  return (
    <>
      <Header />

      {/* Hero section */}
      <div className="mt-18 pt-10 px-4">
        <H1 font="inter" text="Where Simple Stories Find Their Voices" others="md:text-center" />
        <p className="text-accent mt-2 mb-10 max-sm:mb-15 text-sm md:mx-auto md:text-center md:w-[70%]">Clivo is a platform for writers who demands excellence. Craft, Publish, and share your stories with a community that values quality</p>
        <PhotoGrid />
      </div>

      {/* Subscription Section */}
      <div className="mt-20 p-4">
        <H2 font="exo" text="Stay Inspired" others="text-center" />
        <p className="text-sm text-center text-accent mt-2">Join our community of discerning readers and receive carefully selected stories, writing insights, and exclusive content delivered to your inbox.</p>
        <form onSubmit={subscribe} className="mt-4 px-5 relative">
          <input 
            type="email" 
            className="border-1 border-accent p-3 rounded-xl w-full text-sm placeholder:text-accent pr-27 focus:outline-none"
            value={email}
            placeholder="Enter your email"
            onChange={(e) => setEmail(e.target.value)}
            required
          />
          <button className="btn-primary text-sm font-exo absolute right-[30px] top-[8px]">Subscribe</button>
        </form>
      </div>
    </>
  )
}

export default Home