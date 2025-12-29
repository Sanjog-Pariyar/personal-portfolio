"use client";

import { useLastViewedPhoto } from "@/utils/useLastViewedPhoto";
import { useParams } from "next/navigation";
import { useEffect, useRef } from "react";
import Link from "next/link";
import useSWR from "swr";
import Image from "next/image";
import { ImageProps } from "@/utils/types";
// import getBase64ImageUrl from "@/utils/generateBlurPlaceholder";

const fetcher = (url: string) => fetch(url).then((r) => r.json());

export default function ImageGallery() {
  const { photoId } = useParams();
  const [lastViewedPhoto, setLastViewedPhoto] = useLastViewedPhoto();

  const lastViewedPhotoRef = useRef<HTMLAnchorElement>(null);

  useEffect(() => {
    // This effect keeps track of the last viewed photo in the modal to keep the index page in sync when the user navigates back
    if (lastViewedPhoto && !photoId) {
      lastViewedPhotoRef.current.scrollIntoView({ block: "center" });
      setLastViewedPhoto(null);
    }
  }, [photoId, lastViewedPhoto, setLastViewedPhoto]);

  const { data, error, isLoading } = useSWR(
    "http://localhost:8000/image-transform",
    fetcher
  );

  const resources = data?.resources ?? [];
  
  
  

  const images: ImageProps[] = [];

  let i = 0;
  for (const result of resources) {
    images.push({
      id: i,
      height: result.height,
      width: result.width,
      public_id: result.public_id,
      format: result.format,
      secureUrl: result.secure_url
    });
    i++;
  }

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>Error: {error.message}</div>;

  return (
    <>
      {images.map(({ id, secureUrl }) => (
        <Link
          key={id}
          href={`/?photoId=${id}`}
          as={`/p/${id}`}
          ref={id === Number(lastViewedPhoto) ? lastViewedPhotoRef : null}
          shallow
          className="after:content group relative mb-5 block w-full cursor-zoom-in after:pointer-events-none after:absolute after:inset-0 after:rounded-lg after:shadow-highlight"
        >
          <Image
            loading="eager"
            alt="Next.js Conf photo"
            className="transform rounded-lg brightness-90 transition will-change-auto group-hover:brightness-110"
            style={{ transform: "translate3d(0, 0, 0)" }}
            // placeholder="blur"
            // blurDataURL={blurDataUrl}
            src={secureUrl}
            width={720}
            height={480}
            sizes="(max-width: 640px) 100vw,
                  (max-width: 1280px) 50vw,
                  (max-width: 1536px) 33vw,
                  25vw"
          />
        </Link>
      ))}
    </>
  );
}
