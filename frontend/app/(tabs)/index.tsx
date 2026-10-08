import { StyleSheet } from "react-native";

import Carousel from "@/components/Carousel";
import EditScreenInfo from "@/components/EditScreenInfo";
import { Text, View } from "@/components/Themed";
import { Asset } from "expo-asset";
import { useEffect, useState } from "react";

export default function TabOneScreen() {
  const [localUri, setLocalUri] = useState<string | null>(null);

  useEffect(() => {
    async function loadImage() {
      const [asset] = await Asset.loadAsync(
        require("@/assets/images/images.png"),
      );
      setLocalUri(asset.localUri);
    }
    loadImage();
    console.log(localUri);
  }, []);

  const exampleImages = [
    {
      src: { localUri },
      alt: "Image description 1",
    },
    {
      src: "https://example.com/image2.jpg",
      alt: "Image description 2",
    },
    {
      src: "https://example.com/image3.jpg",
      alt: "Image description 3",
    },
    //...추가적인 이미지 객체들...
  ];
  return (
    <View style={styles.container}>
      <Text style={styles.title}>Tab Oned</Text>
      <Carousel data={exampleImages} />
      <View
        style={styles.separator}
        lightColor="#eee"
        darkColor="rgba(255,255,255,0.1)"
      />
      <EditScreenInfo path="app/(tabs)/index.tsx" />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: "center",
    justifyContent: "center",
  },
  title: {
    fontSize: 20,
    fontWeight: "bold",
  },
  separator: {
    marginVertical: 30,
    height: 1,
    width: "80%",
  },
});
