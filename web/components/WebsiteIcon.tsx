import image from "next/image";
import { useState } from "react";
import styled from "styled-components";

const HitImage = styled(image)`
  border-radius: 5px;
`;

export const WebsiteIcon = ({
  src,
  alt,
  size,
}: {
  src: string;
  alt: string;
  size: number;
}) => {
  const [imgFailure, setImgFailure] = useState(false);

  return (
    <>
      {!imgFailure ? (
        <HitImage
          src={src}
          alt={alt}
          width={size}
          height={size}
          onError={() => setImgFailure(true)}
        />
      ) : (
        <span>{new URL(alt).hostname.charAt(0).toUpperCase()}</span>
      )}
    </>
  );
};
