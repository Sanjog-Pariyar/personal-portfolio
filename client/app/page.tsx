import Container from "@/components/container";
import Image from "next/image";

export default function Home() {
  return (
    <>
      <Container>
        <div className="space-y-6">
          <h1 className="text-2xl font-bold">
            Hey, I am a Software developer. I enjoy understanding how things
            work from the inside.
          </h1>
          <p>
            I build reliable, scalable, and clean software. From modern
            frontends to secure backend systems, I turn ideas into working
            products.
          </p>
          <p>Building things that work — fast, secure, and maintainable.</p>
        </div>
      </Container>

      <div className="container max-w-4xl m-auto px-4 mt-20">
        <Image
          // src="/desk.jpg"
          src="/formalphoto.jpg"
          alt="my desk"
          width={1920 / 2}
          height={1280 / 2}
        />
      </div>
    </>
  );
}
