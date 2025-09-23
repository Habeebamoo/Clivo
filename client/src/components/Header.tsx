import logo from "../assets/logo.jpg"

const Header = () => {
  return (
    <header className="py-4 px-6 sm:px-8 flex-between bg-white fixed top-0 left-0 right-0 z-10">
      <nav className="flex-between w-full">
        <div className="flex-start gap-2">
          <img src={logo} className="h-8 rounded-sm" />
          <h1 className="text-2xl font-bold font-inter">Clivo</h1>
        </div>
        <button className="btn-primary text-sm font-exo px-3 py-2">Get Started</button>
      </nav>
    </header>
  )
}

export default Header