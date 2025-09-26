import { useSelector } from "react-redux"
import { H2 } from "../../components/Typo"
import { BiLink, BiPencil } from "react-icons/bi"
import { MdDateRange, MdVerified } from "react-icons/md"

const MyProfile = () => {
  const user = useSelector((state: any) => state.user.user)

  return (
    <main className="grid grid-cols-1 md:grid-cols-2 md:w-[800px] mx-auto ">
      <section>
        {/* Profile Picture */}
        <div>
          <div className="h-40 bg-mutedLight relative flex-center flex-col">
            <H2 font="exo" text="Clivo" color="accent" />
            <p className="text-[12px] text-accentLight mt-1">Where Simple Stories Find Thier Voices</p>

            {/* profile pic */}
            <div className="bg-white p-1 h-26 w-26 rounded-full flex-center absolute left-7 bottom-[-70px]">
              <div className="h-24 w-24 rounded-full overflow-hidden border-1 border-accentLight">
                {user.picture ? (
                  <img src={user.picture} className="h-full w-full object-cover" />
                ) : (
                  <div className="h-full w-full object-cover bg-muted"></div>
                )}
              </div>
            </div>

          </div>
        </div>

        {/* User */}
        <div className="mt-20 px-8">
          {/* Name */}
          <div className="flex-start gap-2">
            <H2 font="exo" text={user.name} />
            {user.verified ? 
              <MdVerified size={18} color="rgba(93, 110, 189, 1)" /> : 
              <p className="text-[12px] underline cursor-pointer">
                Get Verified
              </p>
            }

          </div>

          {/* username */}
          <p className="text-[12px] text-accentLight underline font-inter">{user.username}</p>

          {/* Bio */}
          <p className="font-exo text-sm text-accent mt-4">{user.bio}</p>
          {user.profileLink && 
            <div className="flex-start gap-1 mt-2">
              <BiLink />
              <p className="text-blue-500 text-[12px] underline cursor-pointer">{user.profileLink}</p>
            </div>
          }
          
          {/* follows */}
          <div className="text-sm text-accent font-inter flex-start mt-2 gap-2">
            <p>{user.following} Following</p>
            <div className="h-1 w-1 bg-accentLight rounded-full"></div>
            <p>{user.followers} Followers</p>
          </div>

          {/* others */}
          <div className="mt-3 flex-start gap-2 text-[12px] font-exo text-accent">
            <MdDateRange size={16} />
            <p>Joined {user.createdAt}</p>
          </div>

          <button className="text-sm flex-center gap-2 py-2 px-3 border-1 border-accent rounded-full hover:bg-muted active:bg-muted mt-6">
            <BiPencil />
            <span>Edit Profile</span>
          </button>
        </div>

      </section>
      <section>

      </section>
    </main>
  )
}

export default MyProfile