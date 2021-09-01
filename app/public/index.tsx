import React from 'react';
import NextLink from 'next/link';
import { Link, LinkProps } from '@chakra-ui/react';

const ModifiedLink: React.FC<LinkProps> = (props) => {
  return (
    <NextLink href={props.href}>
      <Link {...props} onClick={() => (document.activeElement as HTMLElement).blur()} />
    </NextLink>
  );
};

export default ModifiedLink;