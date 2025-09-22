import logo from "../assets/logo.png"

const HomeHeader = () => {
  return (
    <header className="py-4 px-8 flex-between bg-white/10 backdrop-blur-md fixed top-0 left-0 right-0 z-10">
      <div className="flex-start gap-1">
        <img src={logo} className="h-8" />
        <h1 className="text-2xl font-bold font-inter">Clivo</h1>
      </div>
      <button className="btn-primary text-sm font-exo px-3 py-2">Get Started</button>
    </header>
  )
}

export default HomeHeader