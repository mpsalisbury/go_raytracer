package raytracer

type Camera struct {
	Width, Height int
}

// Renders a view of the given World onto a canvas of the configured
// size using the configured camera view.
func (cam Camera) Render(s Shape, shader Shader) *Canvas {
	c := MakeCanvas(cam.Width, cam.Height)
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			c.Set(x, y, shader.ColorAt(s, cam.rayForPixel(x, y)))
		}
	}
	return c
}

func (cam Camera) rayForPixel(x, y int) Ray {
	cameraPos := Point{0, 0, -4}
	filmPoint := cam.filmPointForPixel(x, y)
	direction := filmPoint.Minus(cameraPos).Norm()
	return Ray{cameraPos, direction}
}

// pixels [0..hPixels-1, 0..vPixels-1] map onto film of fieldOfView.
// Film is at z=0, width and height set by fieldOfView from camera at origin.
func (cam Camera) filmPointForPixel(x, y int) Point {
	pZ := 0.0
	pixelScale := 4.0 / float64(cam.Width)
	middleX := float64(cam.Width-1) / 2.0
	middleY := float64(cam.Height-1) / 2.0

	pX := (float64(x) - middleX) * pixelScale
	pY := (float64(cam.Height-1-y) - middleY) * pixelScale

	return Point{pX, pY, pZ}
}

/*
// Describes a Camera. Renders a view of a World into a Canvas.
@AutoValue
public abstract class Camera {

  // @param fromP camera location
  // @param toP camera look-at location
  // @param upV camera up direction
  public static Camera create(
      int hPixels, int vPixels, double fieldOfView, Tuple fromP, Tuple toP, Tuple upV) {
    return new AutoValue_Camera(
        hPixels, vPixels, fieldOfView, Matrix.viewTransform(fromP, toP, upV));
  }

  public abstract int hPixels();
  public abstract int vPixels();
  public abstract double fieldOfView();
  public abstract Matrix transform();

  protected Matrix inverseTransform() {
    return transform().invert();
  }

  // Renders a view of the given World onto a canvas of the configured size using the
  // configured camera view.
  public Canvas render(World world) {
    Canvas c = new Canvas(hPixels(), vPixels());
    c.forEachIndex(
        (x, y) -> {
          c.setPixel(x, y, world.colorAt(rayForPixel(x, y)));
        });
    return c;
  }

  private static final Tuple CAMERA_POS = Tuple.point(0, 0, 0);

  // Returns the camera ray that passes from the camera through the pixel at the given coordinates.
  public Ray rayForPixel(int x, int y) {
    Tuple filmPoint = filmPointForPixel(x, y);
    Tuple direction = filmPoint.minus(CAMERA_POS).normalize();

    return Ray.create(inverseTransform().times(CAMERA_POS), inverseTransform().times(direction));
  }

  // Returns how big each pixel should be in x/y-space units.
  @Memoized
  protected double pixelScale() {
    double filmSize = 2.0 * Math.tan(fieldOfView() / 2); // film size at distance 1.
    int maxPixels = Math.max(hPixels(), vPixels());
    double pixelScale = filmSize / (maxPixels - 1);
    return pixelScale;
  }

  // TODO: correctly map to pixel centers
  // pixels [0..hPixels-1, 0..vPixels-1] map onto film of fieldOfView.
  // Film is at z=-1, width and height set by fieldOfView from camera at origin.
  public Tuple filmPointForPixel(int x, int y) {
    // film Z plane.
    final double PZ = -1.0;
    final double pixelScale = pixelScale();
    final double middleX = (hPixels() - 1) / 2.0;
    final double middleY = (vPixels() - 1) / 2.0;

    double pX = ((hPixels() - 1 - x) - middleX) * pixelScale;
    double pY = ((vPixels() - 1 - y) - middleY) * pixelScale;

    return Tuple.point(pX, pY, PZ);
  }
}
*/
